package com.example.kudaki.profile.owner;


import android.app.ProgressDialog;
import android.content.Intent;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import androidx.cardview.widget.CardView;
import androidx.constraintlayout.widget.ConstraintLayout;
import androidx.fragment.app.Fragment;

import com.example.kudaki.R;
import com.example.kudaki.profile.etalase.AddEtalaseActivity;
import com.example.kudaki.profile.etalase.EtalaseActivity;
import com.example.kudaki.transaction.OwnerTransactionActivity;
import com.orhanobut.hawk.Hawk;

import butterknife.BindView;
import butterknife.ButterKnife;

/**
 * A simple {@link Fragment} subclass.
 */
public class OwnerFragment extends Fragment implements OwnerContract.View{
    @BindView(R.id.ownerPending)
    ConstraintLayout pending;
    @BindView(R.id.ownerAccepted)
    ConstraintLayout approved;
    @BindView(R.id.ownerRented)
    ConstraintLayout rented;
    @BindView(R.id.ownerDone)
    ConstraintLayout done;
    @BindView(R.id.ownerEtalase)
    CardView etalase;
    @BindView(R.id.ownerAddItem)
    Button btnAdd;
    @BindView(R.id.ownerPendingBadge)
    ConstraintLayout pendingBadge;
    @BindView(R.id.ownerAcceptedBadge)
    ConstraintLayout approvedBadge;
    @BindView(R.id.ownerRentedBadge)
    ConstraintLayout rentedBadge;
    @BindView(R.id.ownerDoneBadge)
    ConstraintLayout doneBadge;
    @BindView(R.id.ownerPendingNumber)
    TextView pendingNumber;
    @BindView(R.id.ownerAcceptedNumber)
    TextView approvedNumber;
    @BindView(R.id.ownerRentedNumber)
    TextView rentedNumber;
    @BindView(R.id.ownerDoneNumber)
    TextView doneNumber;

    String token;

    OwnerContract.Presenter contractPresenter;
    OwnerPresenter presenter;

    ProgressDialog progressDialog;

    public OwnerFragment() {
        // Required empty public constructor
    }

    @Override
    public View onCreateView(LayoutInflater inflater, ViewGroup container,
                             Bundle savedInstanceState) {
        View view = inflater.inflate(R.layout.fragment_owner, container, false);
        ButterKnife.bind(this, view);

        Hawk.init(view.getContext()).build();

        token = Hawk.get("token");

        presenter = new OwnerPresenter(this, token);

        progressDialog = new ProgressDialog(view.getContext());

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
                Intent intent = new Intent(v.getContext(), OwnerTransactionActivity.class);
                intent.putExtra("status", "PENDING");
                startActivity(intent);
            }
        });

        approved.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), OwnerTransactionActivity.class);
                intent.putExtra("status", "APPROVED");
                startActivity(intent);
            }
        });

        rented.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), OwnerTransactionActivity.class);
                intent.putExtra("status", "RENTED");
                startActivity(intent);
            }
        });

        done.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), OwnerTransactionActivity.class);
                intent.putExtra("status", "DONE");
                startActivity(intent);
            }
        });

        etalase.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(v.getContext(), EtalaseActivity.class);
                startActivity(intent);
            }
        });

        btnAdd.setOnClickListener(v -> {
            Intent add = new Intent(v.getContext(), AddEtalaseActivity.class);
            startActivity(add);
        });
    }

    @Override
    public void showProgress() {
        progressDialog.setMax(100);
        progressDialog.setMessage("Please wait...");
        progressDialog.setTitle("Loading");
        progressDialog.setProgressStyle(ProgressDialog.STYLE_SPINNER);
        progressDialog.show();
    }

    @Override
    public void closeProgress() {
        progressDialog.dismiss();
    }

    @Override
    public void showAddSuccess(String message) {
        Toast.makeText(getContext(), message, Toast.LENGTH_SHORT).show();
        Intent intent = new Intent(getContext(), EtalaseActivity.class);
        startActivity(intent);
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

    @Override
    public void setPresenter(OwnerContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
