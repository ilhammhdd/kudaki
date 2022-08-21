SELECT
  rg.id,
  rg.uuid,
  rg.upvote,
  rg.downvote,
  rg.seen,
  rg.created_at,
  p.full_name,
  u.email,
  m.name
FROM
  kudaki_mountain.recommended_gears rg
  JOIN kudaki_user.users u ON rg.user_uuid = u.uuid
  JOIN kudaki_user.profiles p ON p.user_uuid = u.uuid
  JOIN kudaki_mountain.mountains m ON rg.mountain_uuid = m.uuid
WHERE
  rg.mountain_uuid = ?;
-- -----------------------------------------------------------------
SELECT file_path FROM kudaki_mountain.mountain_files WHERE mountain_uuid = ?;