package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class Owner {
    @SerializedName("full_name")
    @Expose
    private String fullName;
    @SerializedName("email")
    @Expose
    private String email;
    @SerializedName("phone_number")
    @Expose
    private String phoneNumber;
    @SerializedName("total_price")
    @Expose
    private Integer totalPrice;
    @SerializedName("total_item")
    @Expose
    private Integer totalItem;
    @SerializedName("owner_approval_status")
    @Expose
    private String ownerApprovalStatus;
    @SerializedName("owner_order_uuid")
    @Expose
    private String ownerOrderUuid ;
    @SerializedName("cart_items")
    @Expose
    private ArrayList<CartItem> cartItems = null;

    public String getFullName() {
        return fullName;
    }

    public void setFullName(String fullName) {
        this.fullName = fullName;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPhoneNumber() {
        return phoneNumber;
    }

    public void setPhoneNumber(String phoneNumber) {
        this.phoneNumber = phoneNumber;
    }

    public Integer getTotalPrice() {
        return totalPrice;
    }

    public void setTotalPrice(Integer totalPrice) {
        this.totalPrice = totalPrice;
    }

    public Integer getTotalItem() {
        return totalItem;
    }

    public void setTotalItem(Integer totalItem) {
        this.totalItem = totalItem;
    }

    public String getOwnerApprovalStatus() {
        return ownerApprovalStatus;
    }

    public void setOwnerApprovalStatus(String ownerApprovalStatus) {
        this.ownerApprovalStatus = ownerApprovalStatus;
    }

    public String getOwnerOrderUuid() {
        return ownerOrderUuid;
    }

    public void setOwnerOrderUuid(String ownerOrderUuid) {
        this.ownerOrderUuid = ownerOrderUuid;
    }

    public ArrayList<CartItem> getCartItems() {
        return cartItems;
    }

    public void setCartItems(ArrayList<CartItem> cartItems) {
        this.cartItems = cartItems;
    }
}
