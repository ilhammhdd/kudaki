package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.List;

public class RentalResponse {
    @SerializedName("data")
    @Expose
    private RentalData data;

    public RentalData getData() {
        return data;
    }

    public void setData(RentalData data) {
        this.data = data;
    }

    public class RentalData {
        @SerializedName("Total")
        @Expose
        private Integer total;
        @SerializedName("items")
        @Expose
        private List<Item> items = null;

        public Integer getTotal() {
            return total;
        }

        public void setTotal(Integer total) {
            this.total = total;
        }

        public List<Item> getItems() {
            return items;
        }

        public void setItems(List<Item> items) {
            this.items = items;
        }
    }

    public class Item {
        @SerializedName("Id")
        @Expose
        private String id;
        @SerializedName("Properties")
        @Expose
        private Properties properties;

        public String getId() {
            return id;
        }

        public void setId(String id) {
            this.id = id;
        }

        public Properties getProperties() {
            return properties;
        }

        public void setProperties(Properties properties) {
            this.properties = properties;
        }
    }

    public class Properties {
        @SerializedName("item_amount")
        @Expose
        private String itemAmount;
        @SerializedName("item_description")
        @Expose
        private String itemDescription;
        @SerializedName("item_name")
        @Expose
        private String itemName;
        @SerializedName("item_photo")
        @Expose
        private String itemPhoto;
        @SerializedName("item_price")
        @Expose
        private String itemPrice;
        @SerializedName("item_rating")
        @Expose
        private String itemRating;
        @SerializedName("item_unit")
        @Expose
        private String itemUnit;
        @SerializedName("item_uuid")
        @Expose
        private String itemUuid;
        @SerializedName("storefront_uuid")
        @Expose
        private String storefrontUuid;

        public String getItemAmount() {
            return itemAmount;
        }

        public void setItemAmount(String itemAmount) {
            this.itemAmount = itemAmount;
        }

        public String getItemDescription() {
            return itemDescription;
        }

        public void setItemDescription(String itemDescription) {
            this.itemDescription = itemDescription;
        }

        public String getItemName() {
            return itemName;
        }

        public void setItemName(String itemName) {
            this.itemName = itemName;
        }

        public String getItemPhoto() {
            return itemPhoto;
        }

        public void setItemPhoto(String itemPhoto) {
            this.itemPhoto = itemPhoto;
        }

        public String getItemPrice() {
            return itemPrice;
        }

        public void setItemPrice(String itemPrice) {
            this.itemPrice = itemPrice;
        }

        public String getItemRating() {
            return itemRating;
        }

        public void setItemRating(String itemRating) {
            this.itemRating = itemRating;
        }

        public String getItemUnit() {
            return itemUnit;
        }

        public void setItemUnit(String itemUnit) {
            this.itemUnit = itemUnit;
        }

        public String getItemUuid() {
            return itemUuid;
        }

        public void setItemUuid(String itemUuid) {
            this.itemUuid = itemUuid;
        }

        public String getStorefrontUuid() {
            return storefrontUuid;
        }

        public void setStorefrontUuid(String storefrontUuid) {
            this.storefrontUuid = storefrontUuid;
        }
    }
}
