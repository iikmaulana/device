package query

const (
	CreateDevice = `INSERT INTO cleva.cleva_device.dv_device (imei, gpstype_id, gsm_number, desription, organization_id, create_at) 
					 VALUES ($1, $2, $3, $4, $5, $6) returning imei`
	DeleteDevice = `UPDATE cleva.cleva_device.dv_device SET delete_at = now() where imei = $1`
	GetAllDevice = `select coalesce(dv.imei, '')            as imei,
						   coalesce(dv.gpstype_id, 0)       as gpstype_id,
						   coalesce(dv.gsm_number, '')      as gsm_number,
						   coalesce(dv.desription, '')      as desription,
						   coalesce(dv.organization_id, '') as organization_id,
						   coalesce(dv.create_at::text, '') as create_at,
						   coalesce(dv.update_at::text, '') as update_at,
						   coalesce(dv.delete_at::text, '') as delete_at
					from cleva.cleva_device.dv_device dv
					where dv.delete_at is null`
	GetDeviceById = `select coalesce(dv.imei, '')            as imei,
						   coalesce(dv.gpstype_id, 0)       as gpstype_id,
						   coalesce(dv.gsm_number, '')      as gsm_number,
						   coalesce(dv.desription, '')      as desription,
						   coalesce(dv.organization_id, '') as organization_id,
						   coalesce(dv.create_at::text, '') as create_at,
						   coalesce(dv.update_at::text, '') as update_at,
						   coalesce(dv.delete_at::text, '') as delete_at
					from cleva.cleva_device.dv_device dv
					where dv.delete_at is null and dv.imei = $1`

	QueryCheckImeiDevice = `select exists ( select imei from cleva.cleva_device.dv_device where imei = $1 and delete_at is null limit 1) as value`
)
