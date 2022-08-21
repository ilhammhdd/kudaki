package com.example.kudaki.login;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public class LoginContract {
    interface View extends BaseView<Presenter> {
        void showOnLoginSuccess(String message, String token);
        void showOnLoginFailed(String message);
        void showProgress();
        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void doLogin(String email, String Password);
    }
}
