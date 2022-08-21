package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class CartItem {
    @SerializedName("uuid")
    @Expose
    private String uuid;
    @SerializedName("total_items")
    @Expose
    private Integer totalItems;
    @SerializedName("item")
    @Expose
    private CartStoreItem item;

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public Integer getTotalItems() {
        return totalItems;
    }

    public void setTotalItems(Integer totalItems) {
        this.totalItems = totalItems;
    }

    public CartStoreItem getItem() {
        return item;
    }

    public void setItem(CartStoreItem item) {
        this.item = item;
    }
}
