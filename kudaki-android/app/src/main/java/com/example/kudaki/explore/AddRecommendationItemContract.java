package com.example.kudaki.explore;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;

public interface AddRecommendationItemContract {
    interface View extends BaseView<Presenter> {
        void showAddSuccess(String message);
        void showAddFailed(String message);
        void showProgress();
        void closeProgress();
    }

    interface Presenter extends BasePresenter {
        void add(String item, String total);
    }
}
