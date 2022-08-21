package com.example.kudaki.profile.etalase;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface EditEtalaseContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showEditSuccess(String message);
        void showEditFailed(String message);
    }

    interface Presenter extends BasePresenter {
        void update(String uuid, String name, String desc, String price, String duration);
    }
}
