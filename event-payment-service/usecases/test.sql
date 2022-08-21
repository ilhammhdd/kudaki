SELECT
  di.id,
  di.uuid,
  di.kudaki_event_uuid,
  di.amount,
  di.purchase_amount,
  di.transaction_id_merchant,
  di.request_date_time,
  di.currency,
  di.purchase_currency,
  di.session_id,
  di.name,
  di.email,
  di.basket,
  di.status,
  ki.id AS kudaki_event_id,
  ki.uuid AS kudaki_event_uuid,
  ki.name AS kudaki_event_name,
  ki.latitude AS kudaki_event_latitude,
  ki.longitude AS kudaki_event_longitude,
  ki.venue AS kudaki_event_venue,
  ki.description AS kudaki_event_description,
  ki.ad_duration_from AS kudaki_event_ad_duration_from,
  ki.ad_duration_to AS kudaki_event_ad_duration_to,
  ki.duration_from AS kudaki_event_duration_from,
  ki.duration_to AS kudaki_event_duration_to,
  ki.seen AS kudaki_event_seen,
  ki.status AS kudaki_event_status
FROM
  kudaki_event.doku_invoices di
  JOIN kudaki_event.kudaki_events ke ON di.kudaki_event_uuid = ke.uuid;