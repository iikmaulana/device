package query

const (
	CreateHistory = `INSERT INTO cleva.cleva_device.dv_history (id, imei, vehicle_id, description, time, status, organization_id, create_at)
					  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning id`
	DeleteHistory = `UPDATE cleva.cleva_device.dv_history SET delete_at = now() where id = $1`
	GetAllHistory = `select coalesce(dvh.id, '')            as id,
						   coalesce(dvh.imei, '')            as imei,
						   coalesce(dvh.vehicle_id, '')      as vehicle_id,
						   coalesce(dvh.description, '')     as description,
						   coalesce(dvh.time::text, '')      as time,
						   coalesce(dvh.status::text, '')    as status,
						   coalesce(dvh.organization_id, '') as organization_id,
						   coalesce(dvh.create_at::text, '') as create_at,
						   coalesce(dvh.update_at::text, '') as update_at,
						   coalesce(dvh.delete_at::text, '') as delete_at
					from cleva.cleva_device.dv_history dvh
					where dvh.delete_at is null`
	GetHistoryById = `select coalesce(dvh.id, '')            as id,
						   coalesce(dvh.imei, '')            as imei,
						   coalesce(dvh.vehicle_id, '')      as vehicle_id,
						   coalesce(dvh.description, '')     as description,
						   coalesce(dvh.time::text, '')      as time,
						   coalesce(dvh.status::text, '')    as status,
						   coalesce(dvh.organization_id, '') as organization_id,
						   coalesce(dvh.create_at::text, '') as create_at,
						   coalesce(dvh.update_at::text, '') as update_at,
						   coalesce(dvh.delete_at::text, '') as delete_at
					from cleva.cleva_device.dv_history dvh
					where dvh.delete_at is null and dvh.id = $1`
)
