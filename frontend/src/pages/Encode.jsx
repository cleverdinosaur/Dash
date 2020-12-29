import React, { useState } from 'react';
import Label from '../components/common/Label';
import FormSuccess from '../components/FormSuccess';
import FormError from '../components/FormError';
import GradientButton from '../components/common/GradientButton';
import axios from 'axios';
import { IP } from '../ip';

function Encode(props) {
  const [url, setUrl] = useState('');
  const [file, setFile] = useState({});
  const [loginError, setLoginError] = useState();
  const [success, setSuccess] = useState();
  const [processing, setProcessing] = useState(false);
  const [imageurl, setImageUrl] = useState('');

  const handleChange = (event) => {
    const { name, value } = event.target;
    if (name === 'url') {
      setUrl(value);
    }
    if (name === 'file') {
      setLoginError('');
      if (!['image/png', 'image/jpeg'].includes(event.target.files[0].type)) {
        setLoginError('Only image file is allowed.');
      }
      setFile(event.target.files[0]);
      console.log(event.target.files[0]);
    }
  };

  const onSubmitHandler = (event) => {
    event.preventDefault();
    setProcessing(true);
    if (file.name === undefined) {
      setLoginError('Select file.');
      setProcessing(false);
      return;
    }
    if (url === '') {
      setLoginError('Enter url.');
      setProcessing(false);
      return;
    }
    const data = new FormData();
    data.append('file', file);
    data.append('url', url);
    axios
      .post(`${IP}/encode`, data, {})
      .then((res) => {
        setSuccess('Encoded Successfully.');
        console.log(res.statusText);
        setImageUrl(`${IP}${res.data.file.replace('.', '')}`);
        setProcessing(false);
      })
      .catch((e) => {
        console.log(e);
        setProcessing(false);
      });
  };
  return (
    <div>
      <form className="mt-8" onSubmit={onSubmitHandler}>
        {success && <FormSuccess text={success} />}
        {loginError && <FormError text={loginError} />}
        <div>
          <div className="mb-2">
            <div className="mb-1">
              <Label text="URL" />
            </div>
            <input
              name="url"
              type="text"
              onChange={handleChange}
              value={url}
              className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5"
              placeholder="Enter Url"
            />
          </div>
          <div className="mb-2">
            <div className="mb-1">
              <Label text="Image" />
            </div>
            <input
              name="file"
              type="file"
              onChange={handleChange}
              className="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-md focus:outline-none focus:shadow-outline-blue focus:border-blue-300 focus:z-10 sm:text-sm sm:leading-5"
            />
          </div>
          <br />
          <div className="mb-2">
            <GradientButton type="submit" text="Encode" loading={processing} />
          </div>
          {imageurl && (
            <section className="text-center p-2 mb-2 rounded border border-green-600 bg-green-100">
              <p className="text-green-700 font-bold">
                <span className="ml-1">
                  <a href={imageurl}>Image Link</a>
                </span>
              </p>
            </section>
          )}
        </div>
      </form>
    </div>
  );
}

export default Encode;
