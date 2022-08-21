package com.example.kudaki.transaction;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface OwnerDetailTransactionContract {
    interface View extends BaseView<Presenter> {
        void showActionSuccess(String message);
        void showProgress();
        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void approve();
        void disapprove();
        void rented();
        void done();
    }
}
