package com.example.kudaki.profile;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.ProfileData;

public interface EditProfileContract {
    interface View extends BaseView<Presenter>{
        void showProgress();
        void closeProgress();
        void showProfileData(ProfileData data);
        void showEditSuccess(String message);
        void showEditFailed(String message);
    }

    interface Presenter extends BasePresenter{
        void loadProfile();
        void update(String name, String phone);
    }
}
