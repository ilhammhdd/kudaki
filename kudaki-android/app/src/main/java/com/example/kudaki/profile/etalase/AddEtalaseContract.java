package com.example.kudaki.profile.etalase;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface AddEtalaseContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showAddSuccess(String message);
        void showAddFailed(String message);
    }

    interface Presenter extends BasePresenter {
        void addItem(String photo, String name, String desc, String price, String amount);
    }
}
