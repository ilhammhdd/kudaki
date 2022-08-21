package com.example.kudaki.cart;

import com.example.kudaki.BasePresenter;
import com.example.kudaki.BaseView;
import com.example.kudaki.model.response.CartData;

public interface CartContract {
    interface View extends BaseView<Presenter> {
        void showCartItems(CartData data);
        void showCheckoutSuccess(String message);
    }

    interface Presenter extends BasePresenter {
        void loadItems();
        void checkout(String uuid);
    }
}
