package com.example.kudaki.profile.tenant;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface TenantContract {
    interface View extends BaseView<Presenter> {
        void showPending(int number);
        void showApproved(int number);
        void showRented(int number);
        void showDone(int number);
    }

    interface Presenter extends BasePresenter {
        void loadPendingNumber();
        void loadApprovedNumber();
        void loadRentedNumber();
        void loadDoneNumber();
    }
}
