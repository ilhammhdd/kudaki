package com.example.kudaki.forgotpwd;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface ForgotPwdContract {
    interface View extends BaseView<Presenter> {
        void showSendSuccess(String message);

        void showSendFailed(String message);

        void showProgress();

        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void doSendEmail(String email);
    }
}
