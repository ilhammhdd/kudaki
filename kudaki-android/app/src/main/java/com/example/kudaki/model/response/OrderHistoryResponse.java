package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class OrderHistoryResponse {
    @SerializedName("data")
    @Expose
    private OrderHistoryData data;

    public OrderHistoryData getData() {
        return data;
    }

    public void setData(OrderHistoryData data) {
        this.data = data;
    }
}
