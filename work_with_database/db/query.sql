-- name: insertPassportData :one
INSERT INTO passports (id, series, number)
VALUES (@id, @series, @number) RETURNING id;

-- name: getPassportData :one
SELECT * FROM passports p
WHERE p.id = @id;

