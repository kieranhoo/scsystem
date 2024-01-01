package queries

const (
	OfficeData string = `
		select office.id ,office.name,office.phone_number, manager, 
		organization_id as org_id, organization.name as org_name,
		organization.phone_number as org_phone_number, organization.email as org_email,
		organization.head as org_head, organization.website as org_website
		from office join organization on organization_id = organization.id;
	`
)
