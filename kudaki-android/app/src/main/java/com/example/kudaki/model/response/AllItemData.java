package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;

public class AllItemData {
    @SerializedName("items")
    @Expose
    private ArrayList<StoreItem> items = null;

    public ArrayList<StoreItem> getItems() {
        return items;
    }

    public void setItems(ArrayList<StoreItem> items) {
        this.items = items;
    }
}
