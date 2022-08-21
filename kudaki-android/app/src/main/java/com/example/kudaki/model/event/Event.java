package com.example.kudaki.model.event;

public class Event {
    private String name;
    private String photoPath;

    public Event(String name, String photoPath) {
        this.name = name;
        this.photoPath = photoPath;
    }

    public String getPhotoPath() {
        return photoPath;
    }

    public void setPhotoPath(String photoPath) {
        this.photoPath = photoPath;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
