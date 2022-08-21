package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class StoreItem {
    @SerializedName("uuid")
    @Expose
    private String uuid;
    @SerializedName("storefront_uuid")
    @Expose
    private String storefrontUuid;
    @SerializedName("storefront_name")
    @Expose
    private String storefrontName;
    @SerializedName("name")
    @Expose
    private String name;
    @SerializedName("amount")
    @Expose
    private Integer amount;
    @SerializedName("unit")
    @Expose
    private String unit;
    @SerializedName("price")
    @Expose
    private Integer price;
    @SerializedName("price_duration")
    @Expose
    private String priceDuration;
    @SerializedName("description")
    @Expose
    private String description;
    @SerializedName("photo")
    @Expose
    private String photo;
    @SerializedName("rating")
    @Expose
    private Double rating;
    @SerializedName("length")
    @Expose
    private Integer length;
    @SerializedName("width")
    @Expose
    private Integer width;
    @SerializedName("height")
    @Expose
    private Integer height;
    @SerializedName("unit_of_measurement")
    @Expose
    private String unitOfMeasurement;
    @SerializedName("color")
    @Expose
    private String color;
    @SerializedName("created_at")
    @Expose
    private Integer createdAt;

    public StoreItem(String uuid, String storefrontUuid, String name, Integer price, String description, String photo, Double rating, String priceDuration) {
        this.uuid = uuid;
        this.storefrontUuid = storefrontUuid;
        this.name = name;
        this.price = price;
        this.priceDuration = priceDuration;
        this.description = description;
        this.photo = photo;
        this.rating = rating;
    }

    public String getStorefrontName() {
        return storefrontName;
    }

    public void setStorefrontName(String storefrontName) {
        this.storefrontName = storefrontName;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public String getStorefrontUuid() {
        return storefrontUuid;
    }

    public void setStorefrontUuid(String storefrontUuid) {
        this.storefrontUuid = storefrontUuid;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getAmount() {
        return amount;
    }

    public void setAmount(Integer amount) {
        this.amount = amount;
    }

    public String getUnit() {
        return unit;
    }

    public void setUnit(String unit) {
        this.unit = unit;
    }

    public Integer getPrice() {
        return price;
    }

    public void setPrice(Integer price) {
        this.price = price;
    }

    public String getPriceDuration() {
        return priceDuration;
    }

    public void setPriceDuration(String priceDuration) {
        this.priceDuration = priceDuration;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getPhoto() {
        return photo;
    }

    public void setPhoto(String photo) {
        this.photo = photo;
    }

    public Double getRating() {
        return rating;
    }

    public void setRating(Double rating) {
        this.rating = rating;
    }

    public Integer getLength() {
        return length;
    }

    public void setLength(Integer length) {
        this.length = length;
    }

    public Integer getWidth() {
        return width;
    }

    public void setWidth(Integer width) {
        this.width = width;
    }

    public Integer getHeight() {
        return height;
    }

    public void setHeight(Integer height) {
        this.height = height;
    }

    public String getUnitOfMeasurement() {
        return unitOfMeasurement;
    }

    public void setUnitOfMeasurement(String unitOfMeasurement) {
        this.unitOfMeasurement = unitOfMeasurement;
    }

    public String getColor() {
        return color;
    }

    public void setColor(String color) {
        this.color = color;
    }

    public Integer getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Integer createdAt) {
        this.createdAt = createdAt;
    }
}
