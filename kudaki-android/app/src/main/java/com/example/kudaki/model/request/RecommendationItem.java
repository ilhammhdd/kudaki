package com.example.kudaki.model.request;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class RecommendationItem {
    @SerializedName("item_type")
    @Expose
    private String type;
    @SerializedName("total")
    @Expose
    private Integer total;

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Integer getTotal() {
        return total;
    }

    public void setTotal(Integer total) {
        this.total = total;
    }
}
