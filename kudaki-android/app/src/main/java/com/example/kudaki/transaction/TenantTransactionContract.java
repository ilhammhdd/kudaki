package com.example.kudaki.transaction;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.OrderHistoryData;

public interface TenantTransactionContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showOrderHistoryData(OrderHistoryData data);
    }

    interface Presenter extends BasePresenter {
        void loadTransaction(String status);
    }
}
