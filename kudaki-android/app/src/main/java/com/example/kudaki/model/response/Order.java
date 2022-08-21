package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class Order {
    @SerializedName("order_num")
    @Expose
    private String orderNum;
    @SerializedName("status")
    @Expose
    private String status;
    @SerializedName("created_at")
    @Expose
    private Integer createdAt;
    @SerializedName("total_item")
    @Expose
    private Integer totalItem;
    @SerializedName("total_price")
    @Expose
    private Integer totalPrice;
    @SerializedName("owners")
    @Expose
    private ArrayList<Owner> owners = null;

    public Order(String orderNum, String status, Integer createdAt, Integer totalItem, Integer totalPrice, ArrayList<Owner> owners) {
        this.orderNum = orderNum;
        this.status = status;
        this.createdAt = createdAt;
        this.totalItem = totalItem;
        this.totalPrice = totalPrice;
        this.owners = owners;
    }

    public String getOrderNum() {
        return orderNum;
    }

    public void setOrderNum(String orderNum) {
        this.orderNum = orderNum;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Integer getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Integer createdAt) {
        this.createdAt = createdAt;
    }

    public Integer getTotalItem() {
        return totalItem;
    }

    public void setTotalItem(Integer totalItem) {
        this.totalItem = totalItem;
    }

    public ArrayList<Owner> getOwners() {
        return owners;
    }

    public void setOwners(ArrayList<Owner> owners) {
        this.owners = owners;
    }

    public Integer getTotalPrice() {
        return totalPrice;
    }

    public void setTotalPrice(Integer totalPrice) {
        this.totalPrice = totalPrice;
    }
}
