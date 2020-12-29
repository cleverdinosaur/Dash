import React, { useState } from 'react';
import Label from '../components/common/Label';
import FormSuccess from '../components/FormSuccess';
import FormError from '../components/FormError';
import GradientButton from '../components/common/GradientButton';
import axios from 'axios';
import { IP } from '../ip';

function Decode(props) {
  const [file, setFile] = useState();
  const [success, setSuccess] = useState();
  const [error, setError] = useState();
  const [processing, setProcessing] = useState(false);
  const [url, setUrl] = useState();
  const [urlNotFound, setUrlNotFound] = useState('');

  const handleChange = (event) => {
    const { name } = event.target;
    if (name === 'file') {
      setError('');
      if (!['image/png'].includes(event.target.files[0].type)) {
        setError('Invalid file.');
      }
      setFile(event.target.files[0]);
    }
  };

  const onSubmitHandler = (event) => {
    event.preventDefault();
    setProcessing(true);
    setUrl('');
    setUrlNotFound('');
    if (file.name === undefined) {
      setError('Select file.');
      setProcessing(false);
      return;
    }
    const data = new FormData();
    data.append('file', file);
    axios
      .post(`${IP}/decode`, data, {})
      .then((res) => {
        setSuccess('Decoded Successfully.');
        if (res.data.url === '') {
          setUrlNotFound('NOT_FOUND');
        } else {
          setUrl(res.data.url);
        }
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
        {error && <FormError text={error} />}
        <div>
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
            <GradientButton type="submit" text="Decode" loading={processing} />
          </div>
          {url && (
            <section className="text-center p-2 mb-2 rounded border border-green-600 bg-green-100">
              <p className="text-green-700 font-bold">
                <span className="ml-1" style={{ overflowWrap: 'break-word' }}>
                  {url}
                </span>
              </p>
            </section>
          )}
          {urlNotFound === 'NOT_FOUND' && (
            <section className="text-center p-2 mb-2 rounded border border-red-600 bg-red-100">
              <p className="text-xs text-red-500">Not Found</p>
            </section>
          )}
        </div>
      </form>
    </div>
  );
}

export default Decode;
