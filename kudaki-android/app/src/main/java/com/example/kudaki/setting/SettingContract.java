package com.example.kudaki.setting;

import android.content.Context;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface SettingContract {
    interface View extends BaseView<SettingContract.Presenter> {
        void showLogoutSuccess(String message);

    }

    interface Presenter extends BasePresenter {
        void doLogout(Context context);
    }
}
