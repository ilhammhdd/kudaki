package com.example.kudaki.model;

public class MenuProfile {
    int icon;
    String menuName;

    public MenuProfile(int icon, String menuName) {
        this.icon = icon;
        this.menuName = menuName;
    }

    public int getIcon() {
        return icon;
    }

    public String getMenuName() {
        return menuName;
    }

}
