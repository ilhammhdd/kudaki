package com.example.kudaki.profile.etalase;

import android.app.ProgressDialog;
import android.content.Intent;
import android.database.Cursor;
import android.net.Uri;
import android.os.Bundle;
import android.provider.MediaStore;
import android.util.Log;
import android.view.MenuItem;
import android.widget.Button;
import android.widget.EditText;
import android.widget.Toast;

import androidx.annotation.Nullable;
import androidx.appcompat.app.AppCompatActivity;
import androidx.appcompat.widget.Toolbar;
import androidx.core.app.NavUtils;

import com.example.kudaki.R;
import com.orhanobut.hawk.Hawk;

import java.io.File;

import butterknife.BindView;
import butterknife.ButterKnife;

public class AddEtalaseActivity extends AppCompatActivity implements AddEtalaseContract.View {
    @BindView(R.id.etalaseAddToolbar)
    Toolbar toolbar;
    @BindView(R.id.etalaseAddUpload)
    Button upload;
    @BindView(R.id.etalaseAddName)
    EditText name;
    @BindView(R.id.etalaseAddDesc)
    EditText desc;
    @BindView(R.id.etalaseAddPrice)
    EditText price;
    @BindView(R.id.etalaseAddAmount)
    EditText amount;
    @BindView(R.id.etalaseAddSubmit)
    Button submit;

    String token;

    AddEtalaseContract.Presenter contractPresenter;
    AddEtalasePresenter presenter;

    ProgressDialog progressDialog;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_add_etalase);
        ButterKnife.bind(this);

        setSupportActionBar(toolbar);
        getSupportActionBar().setDisplayHomeAsUpEnabled(true);

        Hawk.init(this).build();

        token = Hawk.get("token");

        presenter = new AddEtalasePresenter(this,token);

        progressDialog = new ProgressDialog(this);
    }

    @Override
    protected void onResume() {
        super.onResume();

        upload.setOnClickListener(v -> {
            Intent intent = new Intent(Intent.ACTION_GET_CONTENT);
            intent.addCategory(Intent.CATEGORY_OPENABLE);
            intent.setType("image/*");
            startActivityForResult(Intent.createChooser(intent, "Pilih dari Gallery"), 1);
        });

        submit.setOnClickListener(v -> contractPresenter.addItem(
                "http://store.lendcreative.com/centermenu/wp-content/uploads/sites/3/2017/02/eiger_eiger-tas-laptop-14-inch-andesite-01-hitam_full04.jpg",
                name.getText().toString(),
                desc.getText().toString(),
                price.getText().toString(),
                amount.getText().toString()
        ));
    }

    @Override
    protected void onActivityResult(int requestCode, int resultCode, @Nullable Intent data) {
        super.onActivityResult(requestCode, resultCode, data);
        try {
            if (resultCode == RESULT_OK) {
                if (requestCode == 1) {
                    Uri selectedImageUri = data.getData();

                    selectedImageUri.toString();
                    // Get the path from the Uri
                    final String path = getPathFromURI(selectedImageUri);
                    if (path != null) {
                        File f = new File(path);
                        selectedImageUri = Uri.fromFile(f);
                    }

                    Log.d("IMAGE", "onActivityResult: " + path);
                    Log.d("IMAGE", "onActivityResult: " + selectedImageUri);
                }
            }
        } catch (Exception e) {
            Log.e("FileSelectorActivity", "File select error", e);
        }
    }

    public String getPathFromURI(Uri contentUri) {
        Cursor cursor = getContentResolver().query(contentUri, null, null, null, null);
        if (cursor == null) { // Source is Dropbox or other similar local file path
            return contentUri.getPath();
        } else {
            cursor.moveToFirst();
            int idx = cursor.getColumnIndex(MediaStore.Images.ImageColumns.DATA);
            return cursor.getString(idx);
        }
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        switch (item.getItemId()) {
            case android.R.id.home:
                NavUtils.navigateUpFromSameTask(this);
                return true;
        }
        return super.onOptionsItemSelected(item);
    }

    @Override
    public void showAddSuccess(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
        NavUtils.navigateUpFromSameTask(this);
    }

    @Override
    public void showAddFailed(String message) {
        Toast.makeText(this, message, Toast.LENGTH_SHORT).show();
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
    public void setPresenter(AddEtalaseContract.Presenter presenter) {
        this.contractPresenter = presenter;
    }
}
