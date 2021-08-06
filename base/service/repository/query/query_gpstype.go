package query

const (
	CreateGpsType = `INSERT INTO cleva.cleva_device.dv_gpstype (id, name, description, create_at) VALUES ($1, $2, $3, $4) returning id`
	DeleteGpsType = `UPDATE cleva.cleva_device.dv_gpstype SET delete_at = now() where id = $1`
	GetAllGpsType = `select coalesce(dvg.id, 0)              as id,
						   coalesce(dvg.name, '')            as name,
						   coalesce(dvg.description, '')     as description,
						   coalesce(dvg.create_at::text, '') as create_at,
						   coalesce(dvg.update_at::text, '') as update_at,
						   coalesce(dvg.delete_at::text, '') as delete_at
					from cleva.cleva_device.dv_gpstype dvg
					where dvg.delete_at is null`
	GetGpsTypeById = `select coalesce(dvg.id, 0)              as id,
						   coalesce(dvg.name, '')            as name,
						   coalesce(dvg.description, '')     as description,
						   coalesce(dvg.create_at::text, '') as create_at,
						   coalesce(dvg.update_at::text, '') as update_at,
						   coalesce(dvg.delete_at::text, '') as delete_at
					from cleva.cleva_device.dv_gpstype dvg
					where dvg.delete_at is null and dvg.id = $1`
)
