package member

const (
	sqlCreateMember  = `INSERT INTO member (id,email,password) VALUES ($1,$2,$3) RETURNING id`
	sqlGetAll        = `SELECT id, email, created_on FROM member`
	sqlGetMemberByID = `
		SELECT id, email, created_on
		FROM member
		WHERE id=$1`
	sqlLogin = `
		SELECT id, email, password
		FROM member
		WHERE email=$1`
)
