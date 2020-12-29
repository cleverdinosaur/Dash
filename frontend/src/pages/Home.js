import React, { useState } from 'react';
import Card from '../components/common/Card';
import GradientBar from '../components/common/GradientBar';
import GradientButton from '../components/common/GradientButton';
import Encode from './Encode';
import Decode from './Decode';

const Home = () => {
  const [screen, setScreen] = useState('encode');
  console.log(screen);
  const changeScreenEncode = () => {
    setScreen('encode');
  };
  const changeScreenDecode = () => {
    setScreen('decode');
  };

  return (
    <>
      <section className="w-full sm:w-1/2 h-screen m-auto p-8 sm:pt-10">
        <GradientBar />
        <Card>
          <div className="flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
            <div className="max-w-md w-full">
              <div style={{ display: 'flex', justifyContent: 'space-around' }}>
                <GradientButton text="Encode" onClick={changeScreenEncode} />
                <GradientButton text="Decode" onClick={changeScreenDecode} />
              </div>

              <br />
              <div>{screen === 'encode' && <Encode />}</div>
              <div>{screen === 'decode' && <Decode />}</div>
            </div>
          </div>
        </Card>
      </section>
    </>
  );
};

export default Home;
