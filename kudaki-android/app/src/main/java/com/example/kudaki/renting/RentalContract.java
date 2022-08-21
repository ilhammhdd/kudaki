package com.example.kudaki.renting;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.AllItemData;

public interface RentalContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void displayItems(AllItemData data);
    }

    interface Presenter extends BasePresenter {
        void loadItems();
    }
}
