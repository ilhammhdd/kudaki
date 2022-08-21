package com.example.kudaki.profile;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.AddressData;
import com.example.kudaki.model.response.ProfileData;

public interface ProfileContract {
    interface View extends BaseView<Presenter> {
        void showProfileData(ProfileData data);
        void checkAddress(AddressData data);
        void showProgress();
        void closeProgress();
    }

    interface Presenter extends BasePresenter{
        void loadProfile();
    }
}
