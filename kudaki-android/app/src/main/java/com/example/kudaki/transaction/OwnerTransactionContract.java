package com.example.kudaki.transaction;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.OwnerHistoryData;

public interface OwnerTransactionContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showOrderHistoryData(OwnerHistoryData data);
    }

    interface Presenter extends BasePresenter {
        void loadTransaction(String status);
    }
}
