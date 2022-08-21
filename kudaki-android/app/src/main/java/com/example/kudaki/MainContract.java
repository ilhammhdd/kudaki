package com.example.kudaki;

import com.example.kudaki.model.response.MountainData;

public interface MainContract {
    interface View extends BaseView<Presenter> {
        void showPopularData(MountainData data);
    }

    interface Presenter extends BasePresenter {
        void loadPopular();
    }
}
