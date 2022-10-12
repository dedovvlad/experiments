-- name: insertPassportData :exec
INSERT INTO passports (id, series, number)
VALUES (@id, @series, @number);

-- name: getPassportData :one
SELECT * FROM passports p
WHERE p.id = @id;

