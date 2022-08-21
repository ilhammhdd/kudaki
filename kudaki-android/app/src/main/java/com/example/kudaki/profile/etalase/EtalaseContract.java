package com.example.kudaki.profile.etalase;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.StoreData;

public interface EtalaseContract {
    interface View extends BaseView<Presenter> {
        void showEtalaseItem(StoreData data);
        void showProgress();
        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void loadItems();
    }
}
