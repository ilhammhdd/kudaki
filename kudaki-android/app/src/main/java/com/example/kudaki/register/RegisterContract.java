package com.example.kudaki.register;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.user.User;

public interface RegisterContract {
    interface View extends BaseView<Presenter> {
        void showOnRegisterSuccess(String message);

        void showOnRegisterFailed(String message);

        void showLoginActivity();
        void showProgress();
        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void doRegister(User user);

        void backToLogin();

        boolean validatePassword(String password, String confirmPassword);
    }
}
