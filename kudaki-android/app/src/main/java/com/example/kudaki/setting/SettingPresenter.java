package com.example.kudaki.setting;

import android.content.Context;
import android.content.SharedPreferences;

public class SettingPresenter implements SettingContract.Presenter {

    private SettingContract.View settingView;

    public SettingPresenter(SettingContract.View settingView) {
        this.settingView = settingView;
        this.settingView.setPresenter(this);
    }

    @Override
    public void doLogout(Context context) {
        SharedPreferences spLoginToken = context.getSharedPreferences("LoginToken", Context.MODE_PRIVATE);
        SharedPreferences.Editor edLoginToken = spLoginToken.edit();
        edLoginToken.clear();
        edLoginToken.apply();

        SharedPreferences spUser = context.getSharedPreferences("User", Context.MODE_PRIVATE);
        SharedPreferences.Editor edUser = spUser.edit();
        edUser.clear();
        edUser.apply();

        settingView.showLogoutSuccess("Berhasil Keluar");
    }

    @Override
    public void start() {

    }
}
