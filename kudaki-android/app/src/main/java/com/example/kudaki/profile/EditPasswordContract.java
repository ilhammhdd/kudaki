package com.example.kudaki.profile;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface EditPasswordContract {
    interface View extends BaseView<Presenter>{
        void showProgress();
        void closeProgress();
        void showEditSuccess(String message);
        void showEditFailed(String message);
    }

    interface Presenter extends BasePresenter{
        void update(String oldPwd, String newPwd);
    }
}
