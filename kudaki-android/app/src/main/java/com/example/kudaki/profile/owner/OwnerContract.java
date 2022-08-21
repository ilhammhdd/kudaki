package com.example.kudaki.profile.owner;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface OwnerContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showAddSuccess(String message);

        void showPending(int number);
        void showApproved(int number);
        void showRented(int number);
        void showDone(int number);
    }

    interface Presenter extends BasePresenter {
        void addItem(String name, String desc, String price, String duration);
        void loadPendingNumber();
        void loadApprovedNumber();
        void loadRentedNumber();
        void loadDoneNumber();
    }
}
