package com.example.kudaki.renting;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface DetailEquipmentContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showAddSuccess(String message);
        void showCheckoutSuccess(String message);
    }

    interface Presenter extends BasePresenter {
        void addItem();
        void checkout();
    }
}
