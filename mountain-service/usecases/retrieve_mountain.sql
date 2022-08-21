SELECT
  m.uuid,
  m.name,
  m.height,
  m.latitude,
  m.longitude,
  m.difficulty,
  m.description,
  m.created_at
FROM
  (
    SELECT
      m_i.id
    FROM
      kudaki_mountain.mountains m_i
    LIMIT
      ?, ?
  ) m_ids
  JOIN kudaki_mountain.mountains m ON m_ids.id = m.id;
-- here
SELECT mf.file_path FROM kudaki_mountain.mountain_files mf WHERE mf.mountain_uuid = ?;