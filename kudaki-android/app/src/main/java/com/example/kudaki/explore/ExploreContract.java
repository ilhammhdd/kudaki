package com.example.kudaki.explore;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.MountainData;

public interface ExploreContract {
    interface View extends BaseView<Presenter> {
        void showProgress();
        void closeProgress();
        void showMountainData(MountainData data);
    }

    interface Presenter extends BasePresenter {
        void loadMountain();
    }
}

