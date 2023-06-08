package query

var (
	CreateUser = `
    INSERT INTO
      users(
      username,
      email,
      password
      )VALUES(
      :username,
      :email,
      :password
    );
  `

	FindOneUser = `
    SELECT
    %v
    FROM users
    WHERE %v LIMIT 1
  `
)
