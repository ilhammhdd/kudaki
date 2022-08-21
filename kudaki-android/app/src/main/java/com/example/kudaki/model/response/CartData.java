package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.List;

public class CartData {
    @SerializedName("uuid")
    @Expose
    private String uuid;
    @SerializedName("cart")
    @Expose
    private Cart cart;
    @SerializedName("storefronts")
    @Expose
    private List<Storefront> storefronts = null;

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public Cart getCart() {
        return cart;
    }

    public void setCart(Cart cart) {
        this.cart = cart;
    }

    public List<Storefront> getStorefronts() {
        return storefronts;
    }

    public void setStorefronts(List<Storefront> storefronts) {
        this.storefronts = storefronts;
    }
}
