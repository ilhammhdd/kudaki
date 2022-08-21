package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class OwnerHistoryData {
    @SerializedName("owner_orders")
    @Expose
    private ArrayList<OrderOwner> orders = null;

    public ArrayList<OrderOwner> getOrders() {
        return orders;
    }

    public void setOrders(ArrayList<OrderOwner> orders) {
        this.orders = orders;
    }
}
