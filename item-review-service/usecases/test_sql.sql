SELECT
  i.id,
  i.uuid,
  i.storefront_uuid,
  i.name,
  i.amount,
  i.unit,
  i.price,
  i.price_duration,
  i.description,
  i.photo,
  i.rating,
  i.total_raw_rating,
  i.length,
  i.width,
  i.height,
  i.color,
  i.unit_of_measurement,
  i.created_at,
  (
    SELECT
      COUNT(ir_i.id)
    FROM
      kudaki_store.item_reviews ir_i
      JOIN kudaki_store.items i_i ON ir_i.item_uuid = i_i.uuid
      JOIN kudaki_store.storefronts sf_i ON i_i.storefront_uuid = sf_i.uuid
    WHERE
      sf_i.user_uuid = "f5a08be2-871a-4e89-91ff-9c3ca6776fbf"
  ) counted_reviews
FROM
  kudaki_store.items i
  JOIN kudaki_rental.cart_items ci ON i.uuid = ci.item_uuid
  JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid
WHERE
  ci.cart_uuid =(
    SELECT
      c_i.uuid
    FROM
      kudaki_rental.carts c_i
      JOIN kudaki_order.orders o_i ON c_i.uuid = o_i.cart_uuid
    WHERE
      o_i.uuid = "e5357708-9b5b-49e6-a66e-5d8c285f85fd"
  )
  AND sf.user_uuid = "f5a08be2-871a-4e89-91ff-9c3ca6776fbf";