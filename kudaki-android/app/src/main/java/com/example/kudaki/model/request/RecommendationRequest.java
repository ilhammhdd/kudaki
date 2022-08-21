package com.example.kudaki.model.request;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class RecommendationRequest {
    @SerializedName("mountain_uuid")
    @Expose
    private String uuid;
    @SerializedName("items")
    @Expose
    private ArrayList<RecommendationItem> items = null;

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public ArrayList<RecommendationItem> getItems() {
        return items;
    }

    public void setItems(ArrayList<RecommendationItem> items) {
        this.items = items;
    }
}
