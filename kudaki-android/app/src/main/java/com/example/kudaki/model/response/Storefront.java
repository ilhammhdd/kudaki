package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class Storefront {
    @SerializedName("owner_name")
    @Expose
    private String ownerName;
    @SerializedName("owner_email")
    @Expose
    private String ownerEmail;
    @SerializedName("owner_phone_number")
    @Expose
    private String ownerPhoneNumber;
    @SerializedName("cart_items")
    @Expose
    private ArrayList<CartItem> cartItems = null;

    public Storefront(String ownerName, String ownerEmail, String ownerPhoneNumber, ArrayList<CartItem> cartItems) {
        this.ownerName = ownerName;
        this.ownerEmail = ownerEmail;
        this.ownerPhoneNumber = ownerPhoneNumber;
        this.cartItems = cartItems;
    }

    public String getOwnerName() {
        return ownerName;
    }

    public void setOwnerName(String ownerName) {
        this.ownerName = ownerName;
    }

    public String getOwnerEmail() {
        return ownerEmail;
    }

    public void setOwnerEmail(String ownerEmail) {
        this.ownerEmail = ownerEmail;
    }

    public String getOwnerPhoneNumber() {
        return ownerPhoneNumber;
    }

    public void setOwnerPhoneNumber(String ownerPhoneNumber) {
        this.ownerPhoneNumber = ownerPhoneNumber;
    }

    public ArrayList<CartItem> getCartItems() {
        return cartItems;
    }

    public void setCartItems(ArrayList<CartItem> cartItems) {
        this.cartItems = cartItems;
    }
}
