package com.example.kudaki.model.response;

import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

import java.util.List;

public class ErrorResponse {
    @SerializedName("errors")
    @Expose
    private List<String> errors = null;

    /**
     * No args constructor for use in serialization
     */
    public ErrorResponse() {
    }

    /**
     * @param errors
     */
    public ErrorResponse(List<String> errors) {
        super();
        this.errors = errors;
    }

    public List<String> getErrors() {
        return errors;
    }

    public void setErrors(List<String> errors) {
        this.errors = errors;
    }
}
