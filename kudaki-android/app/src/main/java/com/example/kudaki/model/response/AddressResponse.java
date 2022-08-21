package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class AddressResponse {
    @SerializedName("data")
    @Expose
    private AddressData data;

    public AddressData getData() {
        return data;
    }

    public void setData(AddressData data) {
        this.data = data;
    }
}
