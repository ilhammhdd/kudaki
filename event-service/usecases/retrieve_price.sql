SELECT
  ke.seen,
  ke.name,
  ke.venue,
  ke.description,
  ke.duration_from,
  ke.duration_to,
  ke.ad_duration_from,
  ke.ad_duration_to,
  ke.status,
  ke.file_path
FROM
  kudaki_event.kudaki_events ke
WHERE
  ke.uuid = "a4ebad1f-e84f-4ece-aeb9-7826491e7cb8";