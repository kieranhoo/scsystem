package queries

const (
	HistoryData string = `
	SELECT history.id, activity_type, time, admin_id, registration_time, supervisor, 
		user_id, start_day, end_day, first_name, last_name, users.email, users.phone_number, 
		room.name AS room_name, office.name AS office_name, organization.name AS org_name  
	FROM history
		JOIN registration ON registration_id = registration.id
		JOIN users ON users.id = user_id
		JOIN room ON room.id = room_id
		JOIN office ON office.id = office_id
		JOIN organization ON organization.id = organization_id
	WHERE room.id = ? AND DATE(time) = ?
	ORDER BY history.id DESC;
	`

	RegistrationLatest string = `
	SELECT registration.id AS registration_id,  registration_time, user_id AS student_id, start_day, end_day, first_name, last_name, email, phone_number, org_name, office_name, room_name, supervisor
	FROM registration 
		JOIN users ON users.id = registration.user_id
    	JOIN (
			SELECT organization.name AS org_name, office.name AS office_name, room.name AS room_name, room.id 
			FROM room 
				JOIN office ON room.office_id=office.id 
				JOIN organization ON organization.id=office.organization_id
		) AS rooms ON rooms.id = registration.room_id
	WHERE user_id = ? AND room_id = ?
	ORDER BY registration.id DESC LIMIT 1;
	`

	RoomData string = `
	SELECT
		room.id AS room_id, office_id, description, 
		organization_id, office.name AS office_name, organization.name AS organization_name, room.name AS room_name, 
		address, email, head, website, manager, office.phone_number
	FROM room
	JOIN office ON office.id = office_id
    JOIN organization ON organization.id = organization_id;
	`
)
