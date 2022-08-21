package com.example.kudaki.profile.tenant;


import android.content.Intent;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.TextView;

import androidx.constraintlayout.widget.ConstraintLayout;
import androidx.fragment.app.Fragment;

import com.example.kudaki.R;
import com.example.kudaki.transaction.TenantTransactionActivity;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * A simple {@link Fragment} subclass.
 */
public class TenantFragment extends Fragment implements TenantContract.View {
    @BindView(R.id.tenantPending)
    ConstraintLayout pending;
    @BindView(R.id.tenantApproved)
    ConstraintLayout approved;
    @BindView(R.id.tenantRented)
    ConstraintLayout rented;
    @BindView(R.id.tenantDone)
    ConstraintLayout done;
    @BindView(R.id.tenantPendingBadge)
    ConstraintLayout pendingBadge;
    @BindView(R.id.tenantApprovedBadge)
    ConstraintLayout approvedBadge;
    @BindView(R.id.tenantRentedBadge)
    ConstraintLayout rentedBadge;
    @BindView(R.id.tenantDoneBadge)
    ConstraintLayout doneBadge;
    @BindView(R.id.tenantPendingNumber)
    TextView pendingNumber;
    @BindView(R.id.tenantApprovedNumber)
    TextView approvedNumber;
    @BindView(R.id.tenantRentedNumber)
    TextView rentedNumber;
    @BindView(R.id.tenantDoneNumber)
    TextView doneNumber;

    String token;

    TenantContract.Presenter contractPresenter;
    TenantPresenter presenter;

    public TenantFragment() {
        // Required empty public constructor
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_tenant, container, false);
        ButterKnife.bind(this, view);

        Hawk.init(view.getContext()).build();

        token = Hawk.get("token");

        presenter = new TenantPresenter(this, token);

        return view;
    }

    @Override
    public void onStart() {
        super.onStart();

        pendingBadge.setVisibility(View.INVISIBLE);
        approvedBadge.setVisibility(View.INVISIBLE);
        rentedBadge.setVisibility(View.INVISIBLE);
        doneBadge.setVisibility(View.INVISIBLE);

        contractPresenter.loadPendingNumber();
        contractPresenter.loadApprovedNumber();
        contractPresenter.loadRentedNumber();
        contractPresenter.loadDoneNumber();
    }

    @Override
    public void onResume() {
        super.onResume();

        pending.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), TenantTransactionActivity.class);
                intent.putExtra("status", "PENDING");
                startActivity(intent);
            }
        });

        approved.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), TenantTransactionActivity.class);
                intent.putExtra("status", "APPROVED");
                startActivity(intent);
            }
        });

        rented.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), TenantTransactionActivity.class);
                intent.putExtra("status", "RENTED");
                startActivity(intent);
            }
        });

        done.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), TenantTransactionActivity.class);
                intent.putExtra("status", "DONE");
                startActivity(intent);
            }
        });
    }

    @Override
    public void setPresenter(TenantContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }

    @Override
    public void showPending(int number) {
        if (number != 0) {
            pendingBadge.setVisibility(View.VISIBLE);
            pendingNumber.setText(String.valueOf(number));
        }
    }

    @Override
    public void showApproved(int number) {
        if (number != 0) {
            approvedBadge.setVisibility(View.VISIBLE);
            approvedNumber.setText(String.valueOf(number));
        }
    }

    @Override
    public void showRented(int number) {
        if (number != 0) {
            rentedBadge.setVisibility(View.VISIBLE);
            rentedNumber.setText(String.valueOf(number));
        }
    }

    @Override
    public void showDone(int number) {
        if (number != 0) {
            doneBadge.setVisibility(View.VISIBLE);
            doneNumber.setText(String.valueOf(number));
        }
    }
}
