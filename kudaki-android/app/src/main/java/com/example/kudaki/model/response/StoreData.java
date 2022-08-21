package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.List;

public class StoreData {
    @SerializedName("storefront")
    @Expose
    private StoreStorefront storefront;
    @SerializedName("items")
    @Expose
    private List<StoreItem> items = null;

    public StoreStorefront getStorefront() {
        return storefront;
    }

    public void setStorefront(StoreStorefront storefront) {
        this.storefront = storefront;
    }

    public List<StoreItem> getItems() {
        return items;
    }

    public void setItems(List<StoreItem> items) {
        this.items = items;
    }
}
