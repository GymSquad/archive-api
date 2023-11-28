SELECT
	c.id AS campus_id,
	c.name AS campus_name,
	d.id AS department_id,
	d.name AS department_name,
	o.id AS office_id,
	o.name AS office_name,
	w.id AS website_id,
	w.name AS website_name,
	w.url AS website_url,
	COUNT(*) OVER() AS total_results
FROM 
	"Category" c
JOIN
	"Department" d ON c.id = d."categoryId"
JOIN
	"Office" o ON d.id = o."departmentId"
JOIN
	"_OfficeToWebsite" otw ON o.id = otw."A"
JOIN
	"Website" w ON otw."B" = w.id
WHERE 
	(
		c.id > $1
		OR (c.id = $1 AND d.id > $2)
		OR (c.id = $1 AND d.id = $2 AND o.id > $3)
		OR (c.id = $1 AND d.id = $2 AND o.id = $3 AND w.id > $4)
	)
	AND 
	(
		$5::TEXT IS NULL
		OR c.name ILIKE '%' || $5 || '%'
		OR d.name ILIKE '%' || $5 || '%'
		OR o.name ILIKE '%' || $5 || '%'
		OR w.name ILIKE '%' || $5 || '%'
		OR w.url ILIKE '%' || $5 || '%'
	)
ORDER BY
	c.id, d.id, o.id, w.id
LIMIT
	$6;
