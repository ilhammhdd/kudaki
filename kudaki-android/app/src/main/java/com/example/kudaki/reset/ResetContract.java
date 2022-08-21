package com.example.kudaki.reset;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public class ResetContract {
    interface View extends BaseView<Presenter> {
        void showResetSuccess(String message);

        void showResetFailed(String message);

        void showProgress();

        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void doReset(String token, String newPwd);
    }
}
